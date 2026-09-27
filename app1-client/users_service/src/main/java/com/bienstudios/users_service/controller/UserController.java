package com.bienstudios.users_service.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.bienstudios.users_service.service.UserService;
import com.bienstudios.users_service.model.dto.UserRequest;
import com.bienstudios.users_service.model.dto.UserResponse;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;

// TODO: Inyectar versión por variable de entorno, ó "application.properties"
@RestController
@RequestMapping("/api/v26.00.0000_ALPHA/users")
@RequiredArgsConstructor
public class UserController {

    private final UserService service;

    @PostMapping
    public ResponseEntity<UserResponse> create(
        @Valid
        @RequestBody
        UserRequest request
    ) {
        UserResponse response = service.create(request);
        return ResponseEntity.status(HttpStatus.CREATED).body(response);
    }

    @GetMapping("/{id}")
    @PreAuthorize("hasRole('ADMIN') or #id == authentication.principal.id")
    public ResponseEntity<UserResponse> getById(
        @PathVariable
        Long id
    ) {
        return ResponseEntity.ok(service.getById(id));
    }

}
