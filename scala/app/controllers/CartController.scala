package controllers

import models.{CartEntity, ProductEntity}
import play.api.libs.json._
import play.api.mvc._
import repository.Repository

import javax.inject._


@Singleton
class CartController @Inject()(cc: ControllerComponents, repository: Repository) extends AbstractController(cc) with SimpleCrud {

  case class CreateCartRequest(products: List[Long])

  implicit val createCartFormat: OFormat[CreateCartRequest] = Json.format[CreateCartRequest]
  implicit val productFormat: OFormat[ProductEntity] = Json.format[ProductEntity]
  implicit val cartFormat: OFormat[CartEntity] = Json.format[CartEntity]

  private var carts = repository.carts

  override def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    carts.find(_.id == id) match {
      case Some(cart) => Ok(Json.toJson(cart))
      case None => NotFound(s"Cart with id $id not found")
    }
  }

  override def getAll: Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(carts))
  }

  override def create: Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateCartRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      cart => {
        val id = if (carts.isEmpty) 1 else carts.map(_.id).max + 1
        val newCart = CartEntity(id, cart.products.map(productId => repository.products.find(_.id == productId).getOrElse(ProductEntity(0, "", "", 0.0))))
        carts = carts :+ newCart
        Created(Json.obj(
          "status" -> "success",
          "message" -> "Cart created",
          "location" -> routes.CartController.getById(id).url))
      }
    )
  }

  override def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateCartRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      cart => {
        carts.find(_.id == id) match {
          case Some(_) =>
            val updatedCart = CartEntity(id, cart.products.map(productId => repository.products.find(_.id == productId).getOrElse(ProductEntity(0, "", "", 0.0))))
            carts = carts.filterNot(_.id == id) :+ updatedCart
            Ok(Json.obj("status" -> "success", "message" -> "Cart updated"))
          case None =>
            NotFound(Json.obj("status" -> "error", "message" -> s"Cart with id $id not found"))
        }
      }
    )
  }

  override def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    carts.find(_.id == id) match {
      case Some(_) =>
        carts = carts.filterNot(_.id == id)
        Ok(Json.obj("status" -> "success", "message" -> "Cart deleted"))
      case None =>
        NotFound(Json.obj("status" -> "error", "message" -> s"Cart with id $id not found"))
    }
  }
}
