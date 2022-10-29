package routes

import (
	"errors"
	"net/http"

	"0chaos.eu/lan-info-back/db"
	"0chaos.eu/lan-info-back/models"
	"0chaos.eu/lan-info-back/util/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func insertService(c *fiber.Ctx) error {
	s := db.Begin(c.Context())
	log := logger.New(false) // REMOVE ME
	if service := (models.Service{}); c.BodyParser(&service) == nil {
		service.ID = uuid.New()
		log.Info().Msg("1")
		query := `INSERT INTO service (id, name, icon, dark_supported, online, hostname, port)
      VALUES (:id, :name, :icon, :dark_supported, :online, :hostname, :port);`
		_, err := s.NamedExec(query, &service)
		log.Info().Msg("2")
		if err != nil {
			return err
		}
		log.Info().Msg("3")
		s.Commit()
		return c.Status(http.StatusCreated).JSON(service)
	} else if services := ([]models.Service{}); c.BodyParser(&services) == nil {
		query := `INSERT INTO service (id, name, icon, dark_supported, online, hostname, port)
      VALUES (:id, :name, :icon, :dark_supported, :online, :hostname, :port);`

		_, err := s.NamedExec(query, &services)
		if err != nil {
			return err
		}
		s.Commit()
		return c.Status(http.StatusCreated).JSON(services)
	} else {
		return c.BodyParser(&service)
	}
}

func selectServices(c *fiber.Ctx) error {
	s := db.Begin(c.Context())

	services := []models.Service{}
	query := `SELECT * FROM service;`
	err := s.Select(&services, query)
	if err != nil {
		return err
	}

	s.Commit()
	return c.JSON(services)
}

func selectService(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	s := db.Begin(c.Context())

	service := models.Service{}
	query := `SELECT * FROM service WHERE id=$1;`
	err = s.Select(&service, query, id.String())
	if err != nil {
		return err
	}

	s.Commit()
	return c.JSON(service)
}

func updateService(c *fiber.Ctx) error {
	service := models.Service{}
	err := c.BodyParser(&service)
	if err != nil {
		return err
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	s := db.Begin(c.Context())

	service.ID = id

	query := `SELECT id FROM service WHERE id=$1;`
	rows, err := s.Query(query, service.ID.String())
	if err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("not found")
	}

	query = `UPDATE service SET id=:id, name=:name, icon=:icon, dark_supported=:dark_supported, online=:online, hostname=:hostname, port=:port WHERE id=:id;`
	_, err = s.NamedExec(query, &service)
	if err != nil {
		return err
	}

	s.Commit()
	return c.Status(http.StatusOK).JSON(service)
}

func deleteService(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	s := db.Begin(c.Context())

	services := []models.Service{}
	query := `SELECT * FROM service WHERE id=$1;`
	err = s.Select(&services, query, id.String())
	if err != nil {
		return err
	} else if len(services) < 1 {
		return errors.New("not found")
	}

	query = `DELETE FROM service WHERE id=$1;`
	_, err = s.Exec(query, id.String())
	if err != nil {
		return err
	}

	s.Commit()
	return c.SendStatus(http.StatusNoContent)
}

func deleteServices(c *fiber.Ctx) error {
	s := db.Begin(c.Context())

	query := `DELETE FROM service;`
	_, err := s.Exec(query)
	if err != nil {
		return err
	}

	s.Commit()
	return c.SendStatus(http.StatusNoContent)
}

func SetupServices(g fiber.Router) {
	g.Post("/", insertService)
	g.Get("/", selectServices)
	g.Get("/:id", selectService)
	g.Put("/:id", updateService)
	g.Delete("/", deleteServices)
	g.Delete("/:id", deleteService)
}
