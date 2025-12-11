package com.url_shortner.url_shortener.controller;


import com.url_shortner.url_shortener.dto.URLDTO;
import com.url_shortner.url_shortener.model.Url;
import com.url_shortner.url_shortener.service.URLService;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.List;

@RestController
public class URLController {

    URLService urlService;
    URLController(URLService urlService){
        this.urlService = urlService;
    }

    @GetMapping("/url")
    public List<URLDTO> getUrl(){
        return urlService.getUrls();
    }

    @GetMapping("/{shorturl}")
    public void redirect(@PathVariable String shorturl, HttpServletResponse response) throws IOException {
        Url urlToRedirect =  urlService.getShortenedURL(shorturl);
        IO.println(urlToRedirect.getUrl());
        if(urlToRedirect != null) response.sendRedirect(urlToRedirect.getUrl());
        else response.sendError(HttpServletResponse.SC_NOT_FOUND);
    }

    @PostMapping("/url")
    public Url SendUrl(@RequestBody String url){
       return urlService.addURL(url);
    }
}
