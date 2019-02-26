// Create by Yale 2018/11/13 14:24
package corns


import  "github.com/robfig/cron"

//https://godoc.org/github.com/robfig/cron
/**

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
 */

func Start(spec string, cmd func(c *cron.Cron))  {
	c :=cron.New()
	f := func() {
		cmd(c)
	}
	c.AddFunc(spec,f)
	c.Start()
}

func StartSpec(spec string, cmd func())  {
	c :=cron.New()
	f := func() {
		cmd()
	}
	c.AddFunc(spec,f)
	c.Start()
}