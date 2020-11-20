import csv 
import json 
import schedule 
import time 

#s = {[[10,50],[60,120]],[[20,50],[60,140]}
def timepicker(s,duration):
	selected_times = {}
	time_marker = 0;
	isavail = false
	while(time_marker<2400):
		isavail = false
		for(person in s):
			if (!isFree(person,time_marker)):
				break
		if(isavail):
			selected_times.add(time_marker)
		if(selected_times.length > 4):
			break
	return selected_times
  
  

def make_json(csvFilePath, jsonFilePath): 
      
    # create a dictionary 
    data = {} 
      
    # Open a csv reader called DictReader 
    with open(csvFilePath, encoding='utf-8') as csvf: 
        csvReader = csv.DictReader(csvf) 
          
        # Convert each row into a dictionary  
        # and add it to data 
        for rows in csvReader: 
              
            # Assuming a column named 'No' to 
            # be the primary key 
            key = rows['No'] 
            data[key] = rows 
  
    # Open a json writer, and use the json.dumps()  
    # function to dump data 
    with open(jsonFilePath, 'w', encoding='utf-8') as jsonf: 
        jsonf.write(json.dumps(data, indent=4)) 
          


#person is a set of time slots that correspond to one person, time_marker is the current time we are checking if available in 2400 time format
def isFree(person,time_marker):
	for slot in person:
		if(slot[0] <= time_marker && slot[1] >= time_marker):
			return true
	return false

if __name__ == "__main__":
	# Task scheduling 
	# After every 10mins geeks() is called.  
	schedule.every(10).minutes.do(geeks) 
  
	# After every hour geeks() is called. 
	schedule.every().hour.do(geeks) 
  
	# Every day at 12am or 00:00 time bedtime() is called. 
	schedule.every().day.at("00:00").do(bedtime) 
  
	# After every 5 to 10mins in between run work() 
	schedule.every(5).to(10).minutes.do(work) 
  
	# Every monday good_luck() is called 
	schedule.every().monday.do(good_luck) 
  
	# Every tuesday at 18:00 sudo_placement() is called 
	schedule.every().tuesday.at("18:00").do(sudo_placement) 
  
	# Loop so that the scheduling task 
	# keeps on running all time. 
	while True: 
  
		# Checks whether a scheduled task  
		# is pending to run or not 
		schedule.run_pending() 
		time.sleep(1) 

	# Driver Code 
  
	# Decide the two file paths according to your  
	# computer system 
	csvFilePath = r'Names.csv'
	jsonFilePath = r'Names.json'
  
	# Call the make_json function 
	make_json(csvFilePath, jsonFilePath)