package repository

import models.{CartEntity, CategoryEntity, ProductEntity}

import javax.inject._

@Singleton
class Repository {
  var products: List[ProductEntity] = List(
    ProductEntity(1, "Product 1", "Description 1", 10.0),
    ProductEntity(2, "Product 2", "Description 2", 20.0),
    ProductEntity(3, "Product 3", "Description 3", 30.0)
  )

  var carts: List[CartEntity] = List(
    CartEntity(1, List(products.find(_.id == 1).get)),
    CartEntity(2, List(products.find(_.id == 2).get)),
    CartEntity(3, List(products.find(_.id == 3).get))
  )

  var categories: List[CategoryEntity] = List(
    CategoryEntity(1, "Category 1", "Description 1"),
    CategoryEntity(2, "Category 2", "Description 2"),
    CategoryEntity(3, "Category 3", "Description 3")
  )
}
