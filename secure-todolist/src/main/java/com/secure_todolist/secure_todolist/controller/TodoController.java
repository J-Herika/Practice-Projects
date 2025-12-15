package com.secure_todolist.secure_todolist.controller;

import com.secure_todolist.secure_todolist.dto.TodoDto;
import com.secure_todolist.secure_todolist.model.UserInfo;
import com.secure_todolist.secure_todolist.service.TodoService;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/todos")
public class TodoController {

    enum TodosStages{
        COMPLETED,
        UNCOMPLETED,
        ALL

    }

    private TodoService todoService;

    public TodoController(TodoService todoService){
        this.todoService = todoService;
    }

    @GetMapping
    public List<TodoDto> getTodos(@AuthenticationPrincipal UserInfo user, @RequestParam TodosStages which){

        long userId = user.getId();
        switch (which){
            case ALL -> { return todoService.getTodos(userId); }
            case UNCOMPLETED -> { return todoService.getUncompletedTodos(userId); }
            case COMPLETED -> { return todoService.getCompletedTodos(userId); }
            default -> throw new IllegalArgumentException("Invalid stage");
        }
    }

    @PostMapping
    public TodoDto addTodo(@RequestBody TodoDto newTodo,@AuthenticationPrincipal UserInfo user){
        TodoDto secureTodo = new TodoDto(newTodo.content(), newTodo.isCompleted(), user.getId());
        return todoService.addTodo(secureTodo);
    }

}
