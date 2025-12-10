package com.backend.todo_api.service;

import com.backend.todo_api.dto.TaskDTO;
import com.backend.todo_api.model.Task;
import com.backend.todo_api.repository.TaskRepository;
import org.springframework.stereotype.Service;
import org.springframework.util.PathMatcher;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.List;
import java.util.Optional;

@Service
public class TaskService {

    private final TaskRepository taskRepository;

    TaskService(TaskRepository taskRepository){
        this.taskRepository = taskRepository;
    }

    public List<TaskDTO> getTasks(){
        List<Task> tasks =  taskRepository.findAll();
        return tasks.stream()
                .map(task -> new TaskDTO(task.getId(), task.getDescription(),task.getIsCompleted()))
                .toList();
    }

    public Task addTask(@RequestBody Task newTask){
        if(newTask.getDescription() == null || newTask.getDescription().trim().isEmpty()){
            throw new IllegalArgumentException("Task description cannot be empty!");
        }
        return taskRepository.save(newTask);
    }

    public void checkTask(@PathVariable long id){
//        get the task you want to change. then update it normaly then save it again.
        Task task = taskRepository.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Task not found with ID: " + id));
            task.toggleIsCompleted();
            taskRepository.save(task);

    }

    public void deleteTask(@PathVariable long id){
        taskRepository.deleteById(id);
    }
}
