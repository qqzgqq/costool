<pre>
usage: costool [OPERATION] -url $BucketURL -sd $SecretID -sk $SecretKEY [PARAMETER]
       costool [OPERATION] -url $BucketURL -sds $SSecretID -sks $SSecretKEY [PARAMETER]

OPERATION:
-upload     Transfer the files to the cos bucket
-download   download the files from the cos bucket
-move       Moves files around in buckets,can use rename
-delete     delete the files from the cos bucket

PARAMETER:
-url   bucket_url
-sd    secretid
-sk    secretkey
-sds   receive encrypt secretid
-sks   receive encrypt secretkey
-src   the source full dir and file name on bukect
-dst   the dest full dir and file name on bukect
-bfn   the file full dir and file name on bucket
-enc   encrypt the string will use as encrypt secretid or secretkey
-fdir  the upload file dir and file name in the system

OPERATION EXAMPLE:
costool upload operation eg1:
        costool -upload -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz
costool upload operationeg 2:
        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey`
        costool -upload -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz

costool download operation eg1:
        costool -download -url https://********* -sd ****** -sk ****** -glt
        please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name
        costool -download -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz
costool download operationeg 2:
        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey`
        costool -download -url https://********* -sds $ssecretid -sks $ssecretkey -glt
        please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name
        costool -download -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz

costool move operation eg1:
        costool -move -url https://********* -sd ****** -sk ****** -src APP_BACKUP/test/test.tar.gz -dst APP_BACKUP/test2.tar.gz
costool move operationeg 2:
        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey`
        costool -move -url https://********* -sds $ssecretid -sks $ssecretkey -src APP_BACKUP/test/test.tar.gz -dst APP_BACKUP/test2.tar.gz

costool delete operation eg1:
        costool -delete -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz
costool delete operationeg 2:
        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey`
        costool -delete -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz
</pre>