package com.url_shortner.url_shortener.model;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Url {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    long id;
    String url;
    String shortenedURL = "";

    public Url() {}

    public Url(String url){
        this.url = url;
    }

    public long getId(){ return this.id;}
    public String getUrl(){ return this.url;}
    public void setUrl(String url){ this.url = url;}
    public String getShortenedURL(){ return this.shortenedURL;}
    public void setShortenedURL(String shortenedURL){ this.shortenedURL = shortenedURL;}
}
