package controllers

import "github.com/gin-gonic/gin"

// AllGetter - return all entries
type AllGetter interface {
	GetAll(*gin.Context)
}

// ByIDGetter - return entry by id
type ByIDGetter interface {
	GetByID(*gin.Context)
}

// Creator - create new entry
type Creator interface {
	Create(*gin.Context)
}

// ByIDUpdater - update entry by id
type ByIDUpdater interface {
	UpdateByID(*gin.Context)
}

// ByIDDeleter - delete entry
type ByIDDeleter interface {
	DeleteByID(*gin.Context)
}
