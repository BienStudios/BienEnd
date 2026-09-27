package com.bienstudios.users_service.mapper;

import java.util.List;

import org.springframework.stereotype.Component;

import com.bienstudios.users_service.model.User;
import com.bienstudios.users_service.model.dto.UserRequest;
import com.bienstudios.users_service.model.dto.UserResponse;

@Component
public class UserMapper {

    public UserResponse toResponse(User user) {
        List<String> roleNames = user.getRoles().stream()
            .map(role -> role.getName())
            .toList();
        return new UserResponse(
            user.getId(),
            user.getEmail(),
            user.getFullName(),
            user.getProvider(),
            roleNames
        );
    }

    public User toEntity(UserRequest request) {
        User user = new User();
        user.setEmail(request.email());
        user.setFullName(request.fullName());
        user.setProvider("local");

        return user;
    }
}
