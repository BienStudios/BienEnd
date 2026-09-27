package com.bienstudios.users_service.exceptions;

public class UnexistingRoleException extends RuntimeException {
    public UnexistingRoleException(String roleName) {
        super("El rol \"" + roleName + "\" no existe o no se encuentra definido.");
    }
}