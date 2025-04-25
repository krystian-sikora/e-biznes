package controllers

import play.api.libs.json.JsValue
import play.api.mvc.{Action, AnyContent}

trait SimpleCrud {
  def getById(id: Long): Action[AnyContent]
  def getAll: Action[AnyContent]
  def create: Action[JsValue]
  def update(id: Long): Action[JsValue]
  def delete(id: Long): Action[AnyContent]
}
