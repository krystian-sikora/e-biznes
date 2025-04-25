package controllers

import models.ProductEntity
import repository.Repository

import javax.inject._
import play.api.libs.json._
import play.api.mvc._


@Singleton
class ProductController @Inject()(cc: ControllerComponents, repository: Repository) extends AbstractController(cc) with SimpleCrud {

  case class CreateProductRequest(name: String, description: String, price: Double)
  implicit val productFormat: OFormat[ProductEntity] = Json.format[ProductEntity]
  implicit val createProductFormat: OFormat[CreateProductRequest] = Json.format[CreateProductRequest]

  private var products = repository.products

  override def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(s"Product with id $id not found")
    }
  }

  override def getAll: Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(products))
  }

  override def create: Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateProductRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      product => {
        val id = if (products.isEmpty) 1 else products.map(_.id).max + 1
        val newProduct = ProductEntity(id, product.name, product.description, product.price)
        products = products :+ newProduct
        Created(Json.obj(
          "status" -> "success",
          "message" -> "Product created",
          "location" -> routes.ProductController.getById(id).url))
      }
    )
  }

  override def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateProductRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      product => {
        products.find(_.id == id) match {
          case Some(_) =>
            val updatedProduct = ProductEntity(id, product.name, product.description, product.price)
            products = products.filterNot(_.id == id) :+ updatedProduct
            Ok(Json.obj("status" -> "success", "message" -> "Product updated"))
          case None =>
            NotFound(Json.obj("status" -> "error", "message" -> s"Product with id $id not found"))
        }
      }
    )
  }

  override def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    products.find(_.id == id) match {
      case Some(_) =>
        products = products.filterNot(_.id == id)
        Ok(Json.obj("status" -> "success", "message" -> "Product deleted"))
      case None =>
        NotFound(Json.obj("status" -> "error", "message" -> s"Product with id $id not found"))
    }
  }
}
