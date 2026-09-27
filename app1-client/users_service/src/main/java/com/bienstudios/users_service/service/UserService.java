package com.bienstudios.users_service.service;

import com.bienstudios.users_service.model.dto.UserRequest;
import com.bienstudios.users_service.model.dto.UserResponse;

public interface UserService {
    UserResponse create(UserRequest request);
    UserResponse getById(Long id);
}
