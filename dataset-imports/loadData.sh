

# Blank mongo them import the dataset
mongo mongodb://localhost:27017 <<EOF
 use datasets
 db.dropDatabase();
 use imports
 db.dropDatabase();
EOF
mongoImport -d datasets -c instances mongodb_datasets
mongoImport -d imports -c imports mongodb_imports

# Blank Neo them import the grapg of dataset, heirachies and codelists
brew services stop neo4j
rm -rf /usr/local/Cellar/neo4j/3.3.0/libexec/data
rm -rf /usr/local/Cellar/neo4j/3.4.5/libexec/data
neo4j-admin load --from=graph.dump --database=graph.db --force
brew services start neo4j
