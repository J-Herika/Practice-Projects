package com.secure_todolist.secure_todolist.service;

import com.secure_todolist.secure_todolist.dto.TodoDto;
import com.secure_todolist.secure_todolist.dto.UserDto;
import com.secure_todolist.secure_todolist.model.Todo;
import com.secure_todolist.secure_todolist.model.UserInfo;
import com.secure_todolist.secure_todolist.repository.TodoRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TodoService {

    TodoRepository todoRepository;

    public TodoService(TodoRepository todoRepository){
        this.todoRepository = todoRepository;
    }

    public List<TodoDto> getTodos(long userID){
        List<Todo> todos = todoRepository.findByUserId(userID);
        return todos.stream().map(this::turnTodoInfoToDto).toList();
    }

    public List<TodoDto> getUncompletedTodos(long userID){
        List<Todo> todos = todoRepository.findByUserIdAndIsCompleted(userID,false);
        return todos.stream().map(this::turnTodoInfoToDto).toList();
    }

    public List<TodoDto> getCompletedTodos(long userID){
        List<Todo> todos = todoRepository.findByUserIdAndIsCompleted(userID,true);
        return todos.stream().map(this::turnTodoInfoToDto).toList();
    }

    public TodoDto addTodo(TodoDto todo){
        Todo newTodo = new Todo( todo.userId(), todo.content(),todo.isCompleted());
        Todo returnedTodo = todoRepository.save(newTodo);;
        return turnTodoInfoToDto(returnedTodo);
    }

    private TodoDto turnTodoInfoToDto(Todo todo){
        return new TodoDto(todo.getContent(),todo.isCompleted(), todo.getUserId());
    }
}
