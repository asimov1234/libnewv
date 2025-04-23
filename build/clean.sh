for i in {0..255}; do 
  dir=$(printf "%02x" $i); 
  rm -rf "$dir"; 
done
