package com.secure_todolist.secure_todolist.model;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Data
@NoArgsConstructor
public class Todo {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    private long userId;
    private String content;
    private boolean isCompleted = false;

    public Todo(long userId, String content, Boolean completed) {
        this.userId = userId;
        this.content = content; // Note: your field is named 'context', but DTO says 'content'
        this.isCompleted = isCompleted;
    }
}
