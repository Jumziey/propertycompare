# Design Notes

## database
Want online syncing, looking at amazon dynamodb for add/remove/list directly. Do want to start with simple file list or similar to start with. Useful option to have and a great way to start.

### Start
We start with trying bbolt and take it from there :).

### Design
Where does the implementation of propertdb interface make most sense? Where should the interface be?

### Further development
Break even points for buying rental properties with rent, management cost, repair cost, tax etc. 

## Architecture notes
- `/app`(Entity) business rules
- `/usecase` (UseCases) Application "business" rules
- `/adapter`(Controllers/Adapters) Converts data from usecases to details
- `/repository/{property}` Data layer implementation
- `/bin/{cmd,api,tools,...}`(Details/Outside deps) cli,database, web etc. 
- `/` log/telemetry/(general stuff many applications can use)
