# CimorghVCF
A wrapper on top of tiledbvcf-cli to manage creation and query for VCF files family.
One powerful feature is that you can add samples to the database gradually as you see fit; meaning you don't need to generate one huge database in one single session.
You can find the Ubuntu build in the bin folder.

> **Note**
> `INPUT_DIR` must be the parent directory where your sample files and other input related files reside in and `OUTPUT_DIR` refers to the directory where any output (like the database files) is placed in. **Also these paths must be absolute.**
> On linux you need `sudo` to run commands.

## Create a Database
When creating a database, you can specify specific subfields(like INFO -> AA) to be stored as seperate attribute which results in faster queries on that field.
```
CimorghVCF create -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME
```
With materialized fields:
```
CimorghVCF create -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME -- -a info_AA
```
## Ingest samples into the database
```
CimorghVCF ingest -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME -s INPUT_DIR/samples_file.txt
```
## Stats
This commands shows stats about the underlying array data, not the genomic statistics.
```
CimorghVCF stat -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME
```
## Sample names
list the names of samples ingested into the database.
```
CimorghVCF samples -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME
```
## Export (vcf,bcf,vcf.gz,gvcf,tsv)
To get full help about all the options related to this command pass the `--help` after `--`:
```
CimorghVCF export -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME -- --help
```
Here is an example of exporting from a database, with fields ALT and REF from two specific samples into a tsv file.
```
CimorghVCF export -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME -- --output_format t --tsv-fields ALT,REF \
    --sample-names Codon-WES-266,Codon-WES-277 --regions 1:11111-123456,3:12321342-234234134 \
    --region-file regions.bed > output.tsv
```
