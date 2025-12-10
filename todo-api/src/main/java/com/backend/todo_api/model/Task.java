package com.backend.todo_api.model;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Task{

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    private String description;
    private boolean isCompleted;

    public Task(){}

    public Task(String description){
        this.description = description;
        this.isCompleted = false;
    }

    public long getId() {return id;}
    public void setId(long id){ this.id = id;}
    public String getDescription(){return description;}
    public void setDescription(String description){ this.description = description;}
    public boolean getIsCompleted(){return isCompleted;}
    public void toggleIsCompleted(){ this.isCompleted = !this.isCompleted;}
}
