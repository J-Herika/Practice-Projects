package com.secure_todolist.secure_todolist.controller;

import com.secure_todolist.secure_todolist.dto.UserDto;
import com.secure_todolist.secure_todolist.model.UserInfo;
import com.secure_todolist.secure_todolist.model.AuthRequest;
import com.secure_todolist.secure_todolist.service.JwtService;
import com.secure_todolist.secure_todolist.service.UserService;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/auth")
public class UserController {

    UserService userService;
    JwtService jwtService;
    AuthenticationManager authenticationManager;

    public UserController(UserService userService,JwtService jwtService,AuthenticationManager authenticationManager){
        this.userService = userService;
        this.jwtService = jwtService;
        this.authenticationManager = authenticationManager;
    }


    @PostMapping("/register")
    public UserDto createUser(@RequestBody UserInfo user){
        return userService.createUser(user);
    }

    @GetMapping("/login")
    public String getUser(@RequestBody AuthRequest authRequest){
        Authentication authenticate = authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(authRequest.getUsername(), authRequest.getPassword())
        );

        if(authenticate.isAuthenticated()){
            return jwtService.generateToken(authRequest.getUsername());
        } else {
            throw new RuntimeException("Invalid Access");
        }
    }
}
