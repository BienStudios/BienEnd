package com.bienstudios.enterprise.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.web.SecurityFilterChain;

@Configuration
public class SecurityConfig {
    
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .csrf(csrf -> csrf.disable()) // APIs Rest/BFF no lo necesita
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/public/**").permitAll()
                .anyRequest().authenticated()
            )
            // Habilitar el flujo de autenticación OAuth2 (redirecciones estándar OIDC)
            .oauth2Login(oauth2 -> oauth2
                .defaultSuccessUrl("/api/v26.00.0000/auth/success", true)
            );
        
        return http.build();
    }
}
