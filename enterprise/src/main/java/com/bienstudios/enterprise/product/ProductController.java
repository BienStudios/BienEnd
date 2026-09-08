package com.bienstudios.enterprise.product;

import com.bienstudios.enterprise.product.dto.ProductMapper;
import com.bienstudios.enterprise.product.dto.ProductResponse;
import com.bienstudios.enterprise.product.dto.ProductRequest;

import lombok.RequiredArgsConstructor;

import jakarta.validation.Valid;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.web.PageableDefault;

/**
 * <h3 lang="es">Controlador de datos de {@code Product}</p>
 * 
 * <p lang="es">Responde a las solicitudes desde el cliente
 * para la lectura y modificación de los productos en la
 * base de datos.</p>
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
@RestController
@RequestMapping("/enterprise/v26.00.0000/product")
@RequiredArgsConstructor
public class ProductController {

    private final ProductService service;
    private final ProductMapper mapper;

    @GetMapping
    public ResponseEntity<Page<ProductResponse>> getProducts(
        @PageableDefault(
            page = 0,
            size = 20,
            sort = "lastUpdate",
            direction = Sort.Direction.DESC
        )
        Pageable pageable
    ) {
        Page<Product> productsPage = service.getAll(pageable);

        Page<ProductResponse> responsePage = productsPage.map(mapper::toResponse);
    
        return ResponseEntity.ok(responsePage);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ProductResponse> getById(@PathVariable Long id) {
        Optional<Product> prod = service.findById(id);
        if (prod.isPresent()) {
            return ResponseEntity.ok(mapper.toResponse(prod.get()));
        } else return ResponseEntity.notFound().build();
    }

    @PostMapping
    public ResponseEntity<ProductResponse> createProduct(
        @Valid
        @RequestBody
        ProductRequest request
    ) {
        Product productToCreate = mapper.toEntity(request);
        Product createdProduct = service.create(productToCreate);

        return ResponseEntity
            .status(HttpStatus.CREATED)
            .body(mapper.toResponse(createdProduct));
    }
}
