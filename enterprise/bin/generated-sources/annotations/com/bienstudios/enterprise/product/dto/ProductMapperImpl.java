package com.bienstudios.enterprise.product.dto;

import com.bienstudios.enterprise.product.Category;
import com.bienstudios.enterprise.product.PaymentCondition;
import com.bienstudios.enterprise.product.PaymentMethod;
import com.bienstudios.enterprise.product.Product;
import com.bienstudios.enterprise.product.SubCategory;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import javax.annotation.processing.Generated;
import org.springframework.stereotype.Component;

@Generated(
    value = "org.mapstruct.ap.MappingProcessor",
    date = "2026-09-07T20:04:11-0300",
    comments = "version: 1.5.5.Final, compiler: Eclipse JDT (IDE) 3.46.100.v20260826-1225, environment: Java 21.0.12.1 (Eclipse Adoptium)"
)
@Component
public class ProductMapperImpl implements ProductMapper {

    @Override
    public Product toEntity(ProductRequest request) {
        if ( request == null ) {
            return null;
        }

        Product.ProductBuilder product = Product.builder();

        product.brand( request.brand() );
        Map<Category, List<SubCategory>> map = request.categories();
        if ( map != null ) {
            product.categories( new LinkedHashMap<Category, List<SubCategory>>( map ) );
        }
        product.condition( request.condition() );
        product.description( request.description() );
        product.health( request.health() );
        product.legality( request.legality() );
        product.model( request.model() );
        Map<PaymentCondition, List<PaymentMethod>> map1 = request.paymentConditions();
        if ( map1 != null ) {
            product.paymentConditions( new LinkedHashMap<PaymentCondition, List<PaymentMethod>>( map1 ) );
        }
        product.price( request.price() );

        return product.build();
    }

    @Override
    public ProductResponse toResponse(Product entity) {
        if ( entity == null ) {
            return null;
        }

        ProductResponse.ProductResponseBuilder productResponse = ProductResponse.builder();

        productResponse.brand( entity.getBrand() );
        Map<Category, List<SubCategory>> map = entity.getCategories();
        if ( map != null ) {
            productResponse.categories( new LinkedHashMap<Category, List<SubCategory>>( map ) );
        }
        productResponse.condition( entity.getCondition() );
        productResponse.createdAt( entity.getCreatedAt() );
        productResponse.description( entity.getDescription() );
        productResponse.health( entity.getHealth() );
        productResponse.id( entity.getId() );
        productResponse.lastUpdate( entity.getLastUpdate() );
        productResponse.legality( entity.getLegality() );
        productResponse.model( entity.getModel() );
        Map<PaymentCondition, List<PaymentMethod>> map1 = entity.getPaymentConditions();
        if ( map1 != null ) {
            productResponse.paymentConditions( new LinkedHashMap<PaymentCondition, List<PaymentMethod>>( map1 ) );
        }
        productResponse.price( entity.getPrice() );

        return productResponse.build();
    }
}
