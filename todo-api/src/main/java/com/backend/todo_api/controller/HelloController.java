package com.backend.todo_api.controller;

import com.backend.todo_api.dto.TaskDTO;
import com.backend.todo_api.model.Task;
import com.backend.todo_api.repository.TaskRepository;
import com.backend.todo_api.service.TaskService;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class HelloController {

    TaskService taskService;

    HelloController(TaskService taskService){
        this.taskService = taskService;
    }


    @GetMapping("/")
    public String sayHello(){
        return "Welcome to the backend world!";
    }

    @GetMapping("/tasks")
    public List<TaskDTO> getTasks(){
        return taskService.getTasks();
    }

    @PostMapping("/tasks")
    public Task addTask(@RequestBody Task newTask){
        return taskService.addTask(newTask);
    }

    @PatchMapping("/tasks/{id}/completed")
    public void checkTask(@PathVariable long id){
        taskService.checkTask(id);
    }

    @DeleteMapping("/tasks/{id}")
    public void deleteTask(@PathVariable long id){
       taskService.deleteTask(id);
    }
}
