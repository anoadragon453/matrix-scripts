for i in account device mediaapi syncapi roomserver serverkey federationsender publicroomsapi naffka; do
  sudo -u postgres dropdb dendrite_$i
  sudo -u postgres createdb -O dendrite dendrite_$i
done
