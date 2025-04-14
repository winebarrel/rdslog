package rdslog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

const (
	APIVersion        = "v15"
	EmptyStringSHA256 = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

type Client struct {
	awsCfg     aws.Config
	endpoint   string
	httpClient *http.Client
	*Options
}

func NewClient(ctx context.Context, options *Options) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return nil, err
	}

	epResolver := rds.NewDefaultEndpointResolverV2()
	epParams := rds.EndpointParameters{Region: aws.String(cfg.Region)}
	ep, err := epResolver.ResolveEndpoint(ctx, epParams)

	if err != nil {
		return nil, err
	}

	client := &Client{
		awsCfg:     cfg,
		endpoint:   ep.URI.String(),
		httpClient: http.DefaultClient,
		Options:    options,
	}

	return client, nil
}

func (client *Client) DownloadCompleteLogFile(ctx context.Context, dst io.Writer) error {
	req, err := client.newRequest(
		"/%s/downloadCompleteLogFile/%s/%s",
		APIVersion, client.DBInstanceIdentifier, client.LogFileName)

	if err != nil {
		return err
	}

	err = client.sign(ctx, req)

	if err != nil {
		return err
	}

	res, err := client.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = io.Copy(dst, res.Body)

	return err
}

func (client *Client) newRequest(pathFmt string, args ...any) (*http.Request, error) {
	u, err := url.Parse(client.endpoint)

	if err != nil {
		return nil, err
	}

	u.Path = fmt.Sprintf(pathFmt, args...)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (client *Client) sign(ctx context.Context, req *http.Request) error {
	creds, err := client.awsCfg.Credentials.Retrieve(ctx)

	if err != nil {
		return err
	}

	signer := v4.NewSigner()
	err = signer.SignHTTP(
		ctx, creds, req, EmptyStringSHA256, "rds", client.awsCfg.Region, time.Now())

	return err
}
