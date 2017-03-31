# Puppetmonitor 
The aim of this project is to provide a simple webpage that takes puppet json reports, turns them into go objects, and displays their status.

## Usage
Query the puppetdb's nodes. The command I use is:

	curl 'https://puppetdb.some.site.com:8081/pdb/query/v4/nodes' --tlsv1 --cacert /etc/puppetlabs/puppet/ssl/certs/ca.pem --cert /etc/puppetlabs/puppet/ssl/certs/some.site.com.pem --key /etc/puppetlabs/puppet/ssl/private_keys/some.site.com.pem

Make sure it's named pdbout.json and the server should parse it.

## Todo
- Have the script take a server and some keys and it should be able to query the server itself every once and a while.
- Custom sorting
- CSS
