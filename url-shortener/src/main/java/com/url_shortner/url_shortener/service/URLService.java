package com.url_shortner.url_shortener.service;

import com.url_shortner.url_shortener.dto.URLDTO;
import com.url_shortner.url_shortener.model.Url;
import com.url_shortner.url_shortener.repository.URLRepository;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.List;

@Service
public class URLService {

    @Autowired
    private URLRepository urlRepository;
    @Autowired
    private Base62Encoder base62Encoder;

    public URLService(URLRepository urlRepository,Base62Encoder base62Encoder){
        this.urlRepository = urlRepository;
        this.base62Encoder = base62Encoder;
    }


    @Transactional
    public Url addURL(@RequestBody String url){

        String cleanUrl = url.trim();
        cleanUrl = cleanUrl.replace("\"", "");

        if(!cleanUrl.startsWith("http://") && !cleanUrl.startsWith("https://")){
            throw new IllegalArgumentException("URL must start with http or https: " + url);
        }
        Url newURL = new Url(cleanUrl);
        newURL = urlRepository.save(newURL);
        String encodedUrl = encodeURL(newURL.getId());
        newURL.setShortenedURL(encodedUrl);
        return newURL;
    }

    public String encodeURL(long urlID){
        return base62Encoder.encode(urlID + 10000); // Added 10000 to make the links look cooler
    }

    @Cacheable(value = "urls", key = "#shortenedUrl")
    public Url getShortenedURL(String shortenedUrl){
        return urlRepository.findByshortenedURL(shortenedUrl);
    }

    public List<URLDTO> getUrls(){
        return urlRepository.findAll().stream().map(url -> new URLDTO(url.getId(), url.getUrl(),url.getShortenedURL())).toList();
    }





}
