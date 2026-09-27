package com.bienstudios.users_service.config;

import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.interfaces.RSAPrivateKey;
import java.security.interfaces.RSAPublicKey;
import java.util.UUID;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.nimbusds.jose.jwk.JWKSet;
import com.nimbusds.jose.jwk.RSAKey;

@Configuration
public class JwsConfig {

    @Bean
    public KeyPair keyPair() {

        try {
            KeyPairGenerator gen = KeyPairGenerator.getInstance("RSA");
            gen.initialize(2048);

            return gen.generateKeyPair();
        } catch (Exception e) {
            throw new IllegalStateException("Error al generar el par de claves RSA", e);
        }
    }

    @Bean
    public RSAKey rsaKey(KeyPair keyPair) {

        RSAPublicKey pubKey = (RSAPublicKey) keyPair.getPublic();
        RSAPrivateKey privKey = (RSAPrivateKey) keyPair.getPrivate();

        return new RSAKey.Builder(pubKey)
            .privateKey(privKey)
            .keyID("bienstudios-key-" + UUID.randomUUID().toString())
            .build();
    }

    @Bean
    public JWKSet jwtSet(RSAKey rsaKey) {
        return new JWKSet(rsaKey);
    }
}
