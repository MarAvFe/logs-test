![CI](https://github.com/armory-io/logAggregation/workflows/CI/badge.svg?branch=master)

<h2>Log Aggregation: Print X sorted lists in order.</h2>

Imagine you have any number of servers (1 to 1000+) that generate log files for your distributed app. Each log file can range from 100MB - 512GB in size. They are copied to your machine which contains only 16GB of RAM.

The local directory would look like this:
```bash
/temp/server-bc329xbv.log
/temp/server-cuyew12x.log
```

Our goal is to print the individual lines out to your screen, sorted by timestamp.

A log file stuctured as a `CSV` with the date in ISO 8601 format in the first column and an event in the second column.

Each individual file is already in time order.

As an example, if file `/temp/server-bc329xbv.log` looks like:

    2016-12-20T19:00:45Z, Server A started.
    2016-12-20T19:01:25Z, Server A completed job.
    2016-12-20T19:02:48Z, Server A terminated.

And file `/temp/server-cuyew12x.log` looks like:

    2016-12-20T19:01:16Z, Server B started.
    2016-12-20T19:03:25Z, Server B completed job.
    2016-12-20T19:04:50Z, Server B terminated.

Then our output would be:

    2016-12-20T19:00:45Z, Server A started.
    2016-12-20T19:01:16Z, Server B started.
    2016-12-20T19:01:25Z, Server A completed job.
    2016-12-20T19:02:48Z, Server A terminated.
    2016-12-20T19:03:25Z, Server B completed job.
    2016-12-20T19:04:50Z, Server B terminated.

---

You can run LogSortApplication to generate log files using `LogGenerator` and then print them with `LogPrinter`.
Feel free to use or modify `LogGenerator` to create more and/or longer log files.
The current `LogPrinter` implementation may work, but is too primitive to meet the above requirements.
This is just a starting point so modify this project any way you want.
