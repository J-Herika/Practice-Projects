package com.url_shortner.url_shortener.repository;

import com.url_shortner.url_shortener.model.Url;
import org.springframework.data.jpa.repository.JpaRepository;

public interface URLRepository extends JpaRepository<Url, Long> {
    Url findByshortenedURL( String shortenedURL);
}
