package com.bienstudios.enterprise.product;

import org.springframework.data.domain.Pageable;

import java.util.Optional;

import org.springframework.data.domain.Page;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor 
public class ProductServiceImpl implements ProductService {

    private final ProductRepository repository;

    @Override
    @Transactional(readOnly = true)
    public Page<Product> getAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Override
    @Transactional(readOnly = true)
    public Optional<Product> findById(Long id) {
        return repository.findById(id);
    }

    @Override
    @Transactional
    public Product create(Product product) {
        // Inyección del nombre de la aplicación
        // TODO: adquirir de archivo de configuración
        product.setApplication("this_application");

        // Añadir un "por defecto" (opcional)
        if (product.getCondition() == null)
            product.setCondition(Condition.USED);

        if (product.getHealth() == null)
            product.setHealth(HealthStatus.UNHEALTHY);

        if (product.getLegality() == null)
            product.setLegality(LegalStatus.RESTRICTED);

        return repository.save(product);
    }
}
