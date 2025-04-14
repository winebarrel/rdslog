package rdslog

type Options struct {
	DBInstanceIdentifier string `kong:"arg='',required,help='The customer-assigned name of the DB instance.'"`
	LogFileName          string `kong:"arg='',required,help='The name of the log file.'"`
}
