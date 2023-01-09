# CimorghVCF
A wrapper on top of tiledbvcf-cli to manage creation and query for VCF files family.
One powerful feature is that you can add samples to the database gradually as you see fit; meaning you don't need to generate one huge database in one single session.
You can find the Ubuntu build in the bin folder.

> **Note**
> `INPUT_DIR` must be the parent directory where your sample files and other input related files reside in and `OUTPUT_DIR` refers to the directory where any output (like the database files) is placed in. **Also these paths must be absolute.**

## Create a Database
```
CimorghVCF create -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME
```
## Ingest samples into the database
```
CimorghVCF ingest -i INPUT_DIR -o OUTPUT_DIR -u DATABASE_NAME -s INPUT_DIR/samples_file.txt
```
## Stats
## Sample names
## Export (vcf,bcf,vcf.gz,gvcf)
