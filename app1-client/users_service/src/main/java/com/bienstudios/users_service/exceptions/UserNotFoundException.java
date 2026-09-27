package com.bienstudios.users_service.exceptions;

public class UserNotFoundException extends RuntimeException{
    public UserNotFoundException(Long id) {
        super("No se encontró ningún usuario con el ID \"" + id.toString() + "\".");
    }

    public UserNotFoundException(String email) {
        super("No se encontró ningún usuario con el correo \"" + email + "\".");
    }
}
