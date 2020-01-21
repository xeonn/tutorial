package com.jejavi.socket;

import lombok.*;

@Getter
@Setter
@AllArgsConstructor
public class OutputMessage {

    private String from;
    private String text;
    private String time;
}
