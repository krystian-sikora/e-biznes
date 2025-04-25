package controllers

import models.CategoryEntity

import javax.inject._
import play.api.libs.json.{JsError, JsValue, Json, OFormat}
import play.api.mvc._
import repository.Repository

@Singleton
class CategoryController @Inject()(cc: ControllerComponents, repository: Repository) extends AbstractController(cc) with SimpleCrud {

  case class CreateCategoryRequest(name: String, description: String)
  implicit val createCategoryFormat: OFormat[CreateCategoryRequest] = Json.format[CreateCategoryRequest]
  implicit val categoryFormat: OFormat[CategoryEntity] = Json.format[CategoryEntity]

  private var categories = repository.categories

  override def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(s"Category with id $id not found")
    }
  }

  override def getAll: Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(categories))
  }

  override def create: Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateCategoryRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      category => {
        val id = if (categories.isEmpty) 1 else categories.map(_.id).max + 1
        val newCategory = CategoryEntity(id, category.name, category.description)
        categories = categories :+ newCategory
        Created(Json.obj(
          "status" -> "success",
          "message" -> "Category created",
          "location" -> routes.CategoryController.getById(id).url))
      }
    )
  }

  override def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[CreateCategoryRequest].fold(
      errors => {
        BadRequest(Json.obj("status" -> "error", "message" -> JsError.toJson(errors)))
      },
      category => {
        categories.find(_.id == id) match {
          case Some(_) =>
            val updatedCategory = CategoryEntity(id, category.name, category.description)
            categories = categories.filterNot(_.id == id) :+ updatedCategory
            Ok(Json.obj("status" -> "success", "message" -> "Category updated"))
          case None =>
            NotFound(s"Category with id $id not found")
        }
      }
    )
  }

  override def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    categories.find(_.id == id) match {
      case Some(_) =>
        categories = categories.filterNot(_.id == id)
        Ok(Json.obj("status" -> "success", "message" -> "Category deleted"))
      case None =>
        NotFound(Json.obj("status" -> "error", "message" -> s"Category with id $id not found"))
    }
  }
}
