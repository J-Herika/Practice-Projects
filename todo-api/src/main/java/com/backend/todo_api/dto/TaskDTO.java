package com.backend.todo_api.dto;

public record TaskDTO(Long id, String description, boolean isCompleted) {
}
