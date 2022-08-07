# logs-test

## Screenshot of execution

![Screenshot from 2022-08-06 22-37-16](https://user-images.githubusercontent.com/8484790/183275885-880b4e33-f147-44cb-9072-e47bce89fcc7.png)

## How does it work?

Considering 1-1000+ files might be read and files may be 100MB-512GB in size, I opted to read all files in chunks, no full file is read, only descriptors are fully tracked into memory.

To avoid memory overflow, I setup a memory limit (e.g. 1GB) which I then divide in the amount of log files in the `logs` folder and divide in 2 because chunk lines will be parsed and loaded into memory as well.

Once limits are in place, the software iterates through all files, reads a chunk size of the file (e.g. 300 chars) and parses as many lines present in said chunk. Then moves to the next file.

After all chunk lines have been read and parsed, it sorts the lines by date, prints them and reads the next set of chunks from the files.

Then all process repeats until all files return a EOF error.

## Personal comments on the solution

1. First of all, after about 3h into the solution, I noticed a big logic error which wasn't solved: The logs from file A, chunk 2, may be dated earlier than those in file B, chunk 1. rendering the solution incorrect. Example below.
2. Memory limit is dynamic according to the amount of log files to aggregate
3. This solution is not useful for live aggregation

## Logic flaw:

Consider these files:
```log
2016-12-21, Server A started.
2016-12-21, Server A completed job.
2016-12-21, Server A terminated.
```

```log
2016-12-20, Server B started.
2016-12-22, Server B completed job.
2016-12-23, Server B terminated.
```

If I read 2 line chunks, the output would be:

```log
2016-12-20, Server B started.
2016-12-21, Server A started.
2016-12-21, Server A completed job.
2016-12-22, Server B completed job.  <-- end of the 1st chunk batch
2016-12-21, Server A terminated.     <-- unsorted log line
2016-12-23, Server B terminated.
```
