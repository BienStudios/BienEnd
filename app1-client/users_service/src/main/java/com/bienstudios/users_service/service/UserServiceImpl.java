package com.bienstudios.users_service.service;

import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import com.bienstudios.users_service.exceptions.UnexistingRoleException;
import com.bienstudios.users_service.exceptions.UserAlreadyExistsException;
import com.bienstudios.users_service.exceptions.UserNotFoundException;
import com.bienstudios.users_service.mapper.UserMapper;
import com.bienstudios.users_service.model.Role;
import com.bienstudios.users_service.model.User;
import com.bienstudios.users_service.model.dto.UserRequest;
import com.bienstudios.users_service.model.dto.UserResponse;
import com.bienstudios.users_service.repository.RoleRepository;
import com.bienstudios.users_service.repository.UserRepository;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor 
public class UserServiceImpl implements UserService {

    private final UserRepository userRepo;
    private final RoleRepository roleRepo;
    private final UserMapper mapper;
    private final PasswordEncoder encoder;

    @Override
    public UserResponse create(UserRequest request) {

        if (userRepo.existsByEmail(request.email()))
            throw new UserAlreadyExistsException(request.email());

        User user = mapper.toEntity(request);
        user.setPasswordHash(encoder.encode(request.password()));

        String roleName = "CUSTOMER";
        Role defaultRole = roleRepo.findByName(roleName)
            .orElseThrow(() -> new UnexistingRoleException(roleName));

        user.getRoles().add(defaultRole);

        return mapper.toResponse(userRepo.save(user));

    }

    @Override
    public UserResponse getById(Long id) {
        User user = userRepo.findById(id)
            .orElseThrow(() -> new UserNotFoundException(id));

        return mapper.toResponse(user);
    }
}
