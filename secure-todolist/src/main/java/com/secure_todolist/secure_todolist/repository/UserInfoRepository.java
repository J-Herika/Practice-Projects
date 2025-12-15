package com.secure_todolist.secure_todolist.repository;

import com.secure_todolist.secure_todolist.model.UserInfo;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface UserInfoRepository extends JpaRepository<UserInfo,Long> {
    Optional<UserInfo> findByName(String name);
}
