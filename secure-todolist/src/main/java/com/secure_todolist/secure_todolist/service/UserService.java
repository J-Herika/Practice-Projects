package com.secure_todolist.secure_todolist.service;

import com.secure_todolist.secure_todolist.dto.TodoDto;
import com.secure_todolist.secure_todolist.dto.UserDto;
import com.secure_todolist.secure_todolist.model.UserInfo;
import com.secure_todolist.secure_todolist.repository.UserInfoRepository;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class UserService {

    UserInfoRepository userInfoRepository;
    PasswordEncoder passwordEncoder;

    public UserService(UserInfoRepository userInfoRepository,PasswordEncoder passwordEncoder){
        this.userInfoRepository = userInfoRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public UserDto createUser(UserInfo newUser){
        if(userInfoRepository.findByName(newUser.getName()).isPresent()) throw new IllegalArgumentException("User already signed up. try to login.");

        newUser.setPassword(passwordEncoder.encode(newUser.getPassword()));
        UserInfo savedUser = userInfoRepository.save(newUser);
        return turnUserInfoToDto(savedUser);
    }

    public UserDto getUser(long userId){
        UserInfo user = userInfoRepository.findById(userId).orElseThrow(() -> new IllegalArgumentException("User not signed up. try to sign up."));
        return turnUserInfoToDto(user);

    }

    private UserDto turnUserInfoToDto(UserInfo user){
        return new UserDto(user.getId(), user.getName(), user.getRole());
    }
}
