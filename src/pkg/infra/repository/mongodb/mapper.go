package mongodb

import "90poe/src/pkg/domain"

func fromDomain(port domain.Port) Port {
	var mongoPort Port
	mongoPort.Key = port.Key
	mongoPort.Name = port.Name
	mongoPort.City = port.City
	mongoPort.Country = port.Country
	mongoPort.Alias = port.Alias
	mongoPort.Regions = port.Regions
	mongoPort.Coordinates = port.Coordinates
	mongoPort.Province = port.Province
	mongoPort.Timezone = port.Timezone
	mongoPort.Unlocs = port.Unlocs
	mongoPort.Code = port.Code
	return mongoPort
}

func (p Port) toDomain() domain.Port {
	var port domain.Port
	port.Key = p.Key
	port.Name = p.Name
	port.City = p.City
	port.Country = p.Country
	port.Alias = p.Alias
	port.Regions = p.Regions
	port.Coordinates = p.Coordinates
	port.Province = p.Province
	port.Timezone = p.Timezone
	port.Unlocs = p.Unlocs
	port.Code = p.Code
	return port
}
