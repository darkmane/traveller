export PROJECT=darkmane-showcase
export REGION=us-central1
export ZONE=$REGION-c
gcloud beta container --project $PROJECT clusters create "cluster-1" \
    --zone "$ZONE" --no-enable-basic-auth --release-channel "regular" \
    --machine-type "n1-standard-1" --image-type "UBUNTU" --disk-type "pd-standard" --disk-size "100" \
    --metadata disable-legacy-endpoints=true \
    --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
    --preemptible --num-nodes "3" --enable-stackdriver-kubernetes --enable-ip-alias \
    --network "projects/$PROJECT/global/networks/default" --subnetwork "projects/$PROJECT/regions/$REGION/subnetworks/default" \
    --default-max-pods-per-node "110" --enable-autoscaling --min-nodes "0" --max-nodes "3" --no-enable-master-authorized-networks \
    --addons HorizontalPodAutoscaling,HttpLoadBalancing --enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0