package com.armory.logsort;

import org.junit.jupiter.api.Test;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class LogPrinterTest {

    @Test
    public void printLogFiles() throws Exception {
        Path tempFile1 = Files.createTempFile("server-bc329xbv", ".log");
        Files.write(tempFile1, (
                "    2016-12-20T19:00:45Z, Server A started.\n" +
                "    2016-12-20T19:01:25Z, Server A completed job.\n" +
                "    2016-12-20T19:02:48Z, Server A terminated.").getBytes());
        Path tempFile2 = Files.createTempFile("server-cuyew12x", ".log");
        Files.write(tempFile2, (
                "    2016-12-20T19:01:16Z, Server B started.\n" +
                "    2016-12-20T19:03:25Z, Server B completed job.\n" +
                "    2016-12-20T19:04:50Z, Server B terminated.").getBytes());
        List<String> logs = new LogPrinter().sortLogs(Arrays.asList(tempFile1.toFile(), tempFile2.toFile()));

        List<String> expected = Arrays.asList(
                "    2016-12-20T19:00:45Z, Server A started.",
                "    2016-12-20T19:01:16Z, Server B started.",
                "    2016-12-20T19:01:25Z, Server A completed job.",
                "    2016-12-20T19:02:48Z, Server A terminated.",
                "    2016-12-20T19:03:25Z, Server B completed job.",
                "    2016-12-20T19:04:50Z, Server B terminated."
        );
        assertEquals(expected, logs);
    }

}