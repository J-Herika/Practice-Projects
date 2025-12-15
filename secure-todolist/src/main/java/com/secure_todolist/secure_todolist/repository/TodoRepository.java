package com.secure_todolist.secure_todolist.repository;

import com.secure_todolist.secure_todolist.model.Todo;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface TodoRepository extends JpaRepository<Todo, Long> {
    List<Todo> findByUserId(long userId);
    List<Todo> findByIsCompleted(boolean isCompleted);
    List<Todo> findByUserIdAndIsCompleted(long userId, boolean isCompleted);

}
