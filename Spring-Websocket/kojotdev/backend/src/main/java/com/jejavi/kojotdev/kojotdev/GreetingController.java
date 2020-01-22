package com.jejavi.kojotdev.kojotdev;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.builder.ToStringBuilder;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.util.HtmlUtils;

@Slf4j
@CrossOrigin
@Controller
public class GreetingController {
    
    @MessageMapping("/hello")
    @SendTo("/topic/greetings")
    public Greeting greeting(HelloMessage message) throws Exception {
        log.info("Received request {}", ToStringBuilder.reflectionToString(message));
        log.info("Delaying 2 seconds ...");
        Thread.sleep(2000); // simulated delay
        log.info("Returning Greeting object");
        return new Greeting("Hello, " + HtmlUtils.htmlEscape(message.getName()));
    }
}
