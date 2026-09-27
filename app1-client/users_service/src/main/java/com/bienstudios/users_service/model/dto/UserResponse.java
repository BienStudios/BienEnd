package com.bienstudios.users_service.model.dto;

import java.util.List;

public record UserResponse(
    Long id,
    String email,
    String fullName,
    String provider,
    List<String> roles
) {}
