package data

const flavorCreate = `
	CREATE TABLE IF NOT EXISTS Flavor (
		FlavorID	INTEGER NOT NULL,
		FlavorName	TEXT,
		PRIMARY KEY(FlavorID)
	)
`

const sizeCreate = `
	CREATE TABLE IF NOT EXISTS Size (
		SizeID	INTEGER NOT NULL,
		SizeName	TEXT,
		PRIMARY KEY(SizeID)
	)
`

const modelCreate = `
	CREATE TABLE IF NOT EXISTS Model (
		ModelID INTEGER NOT NULL,
		ModelName TEXT,
		PRIMARY KEY(ModelID)
	)
`

const coffeeCreate = `
	CREATE TABLE IF NOT EXISTS CoffeeMachine (
		CoffeeMachineID	INTEGER NOT NULL,
		SizeID	INTEGER,
		SKU	TEXT,
		ModelID	INTEGER,
		WaterLine	INTEGER,
		PRIMARY KEY(CoffeeMachineID),
		FOREIGN KEY(SizeID) REFERENCES Size(SizeID),
		FOREIGN KEY(ModelID) REFERENCES Model(ModelID)
	)
`

const podCreate = `
	CREATE TABLE IF NOT EXISTS Pod (
		PodID	INTEGER NOT NULL,
		SizeID	INTEGER,
		FlavorID	INTEGER,
		SKU	TEXT,
		Quantity	INTEGER,
		PRIMARY KEY(PodID),
		FOREIGN KEY(FlavorID) REFERENCES Flavor(FlavorID),
		FOREIGN KEY(SizeID) REFERENCES Size(SizeID)
	)
`

const coffeeMachineQ = `
	SELECT CoffeeMachineID, CoffeeMachine.SizeID, Size.SizeName,
	SKU, CoffeeMachine.ModelID, Model.ModelName, WaterLine FROM CoffeeMachine
	JOIN Size ON Size.SizeID = CoffeeMachine.SizeID
	JOIN Model ON Model.ModelID = CoffeeMachine.ModelID
`

const podQ = `
	SELECT PodID,  Pod.FlavorID, Flavor.FlavorName, Pod.SizeID, Size.SizeName,
	Pod.SKU, Quantity FROM Pod JOIN Flavor ON Flavor.FlavorID = Pod.FlavorID 
	JOIN Size ON Size.SizeID = Pod.SizeID
`

// machineCrossQ takes the pod SizeID parameter to get the proper cross sell.
const machineCrossQ = `
	SELECT PodID,  Pod.FlavorID, Flavor.FlavorName, Pod.SizeID, Size.SizeName,
	Pod.SKU, MIN(Quantity) FROM Pod JOIN Flavor ON Flavor.FlavorID = Pod.FlavorID
	JOIN Size ON Size.SizeID = Pod.SizeID WHERE Pod.SizeID = ? GROUP BY POD.FlavorID
`
