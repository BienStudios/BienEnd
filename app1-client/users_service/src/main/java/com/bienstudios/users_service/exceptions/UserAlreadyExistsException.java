package com.bienstudios.users_service.exceptions;

public class UserAlreadyExistsException extends RuntimeException {
    public UserAlreadyExistsException(String email) {
        super("El correo \"" + email + "\" ya se encuentra registrado.");
    }
}
