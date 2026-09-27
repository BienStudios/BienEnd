package com.bienstudios.users_service.model.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;

public record UserRequest(
    @NotBlank
    @Email
    String email,

    @NotBlank
    String fullName,

    String password // Aplicable únicamente si provider = "local"
) {}
