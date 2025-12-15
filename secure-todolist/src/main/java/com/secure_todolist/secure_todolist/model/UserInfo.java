package com.secure_todolist.secure_todolist.model;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Collection;
import java.util.List;

@Entity
@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserInfo implements UserDetails {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    @Column(unique = true)
    private String name;
    private String password;
    private String role;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return List.of(new SimpleGrantedAuthority(role));
    }

    @Override
    public String getUsername() {return name;}

    @Override
    public boolean isAccountNonExpired() {return UserDetails.super.isAccountNonExpired();}
    @Override
    public boolean isAccountNonLocked() {return UserDetails.super.isAccountNonLocked();}
    @Override
    public boolean isCredentialsNonExpired() {return UserDetails.super.isCredentialsNonExpired();}
    @Override
    public boolean isEnabled() {return UserDetails.super.isEnabled();}
}
