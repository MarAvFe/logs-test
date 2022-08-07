package com.armory.logsort;

import java.io.File;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class LogPrinter {

    public void printLogs(List<File> logFiles) {
        sortLogs(logFiles).forEach(logLine -> System.out.println(logLine));
    }

    protected List<String> sortLogs(List<File> logFiles) {
        List<String> logLines = new ArrayList<>();
        try {
            for (File file : logFiles) {
                logLines.addAll(Files.readAllLines(file.toPath()));
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        Collections.sort(logLines);
        return logLines;
    }

}
