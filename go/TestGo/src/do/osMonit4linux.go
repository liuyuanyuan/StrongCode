package do

import (
	"time"
    //"runtime"
    //"syscall"
    
    "strings"
    "strconv"
    "io/ioutil"
    "errors"
)
 
func GetOSMemory() OSMem {
	filename := "/proc/meminfo"
	lines, _ := readLines(filename)

	mem := OSMem{}
    mem.Current = time.Now().Format(timeFormat)
    
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)

		t, err := strconv.ParseUint(value, 10, 64)
		checkErr(err)
		switch key {
		case "MemTotal":
			mem.Total = t
		case "MemFree":
			mem.Free = t
		/*
		case "Buffers":
			mem.Buffers = t
		case "Cached":
			mem.Cached = t
		case "Active":
			mem.Active = t
		case "Inactive":
			mem.Inactive = t
		*/
		}
	}
	mem.Used = mem.Total - mem.Free
	//mem.Available = mem.Free + mem.Buffers + mem.Cached	
	//mem.UsedPercent = float64(mem.Total-mem.Available) / float64(mem.Total) * 100.0

    Info.Println("Return:(KB)Mem all=", mem.Total, " free=", mem.Free, " used=", mem.Used)
	return mem
}

func GetOSLoadAvg()(OSLoadAvg, error){
	loadAvg := OSLoadAvg{}
	loadAvg.Current = time.Now().Format(timeFormat)
	
	filename := "/proc/loadavg"
	line, err := ioutil.ReadFile(filename)
	if err != nil {
		return loadAvg, err
	}
	
	values := strings.Fields(string(line))
	load1, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return loadAvg, err
	}
	load5, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return loadAvg, err
	}
	load15, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return loadAvg, err
	}

	loadAvg.Load1 = load1
	loadAvg.Load5 = load5
	loadAvg.Load15 = load15
	
	Info.Println("Return: loadAvg load1=", load1, " load5=", load5, "load15=", load15)
	return loadAvg, nil
}

func GetOSCpu() error {
	stat := OSCpu{}
	stat.Current = time.Now().Format(timeFormat)
	
	filename := "/proc/stat"
	var lines = []string{}
	lines, _ = readLinesOffsetN(filename, 0, 1)    
	
	fields := strings.Fields(lines[0])

	if strings.HasPrefix(fields[0], "cpu") == false {
		return errors.New("not contain cpu")
	}

	cpu := fields[0]
	if cpu == "cpu" {
		cpu = "cpu-total"
	}
	user, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return err
	}
	nice, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return err
	}
	system, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return err
	}
	idle, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return err
	}
	iowait, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		return err
	}
	irq, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return err
	}
	softirq, err := strconv.ParseFloat(fields[7], 64)
	if err != nil {
		return err
	}
	stolen, err := strconv.ParseFloat(fields[8], 64)
	if err != nil {
		return err
	}

	cpu_tick := float64(100) // TODO: how to get _SC_CLK_TCK ?

	stat.CPU = cpu
	stat.User = float64(user) / cpu_tick
	stat.Nice = float64(nice) / cpu_tick
	stat.System = float64(system) / cpu_tick
	stat.Idle = float64(idle) / cpu_tick
	stat.Iowait = float64(iowait) / cpu_tick
	stat.Irq = float64(irq) / cpu_tick
	stat.Softirq = float64(softirq) / cpu_tick
	stat.Stolen = float64(stolen) / cpu_tick

	if len(fields) > 9 { // Linux >= 2.6.11
		steal, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			return err
		}
		stat.Steal = float64(steal)
	}
	if len(fields) > 10 { // Linux >= 2.6.24
		guest, err := strconv.ParseFloat(fields[10], 64)
		if err != nil {
			return err
		}
		stat.Guest = float64(guest)
	}
	if len(fields) > 11 { // Linux >= 3.2.0
		guestNice, err := strconv.ParseFloat(fields[11], 64)
		if err != nil {
			return err
		}
		stat.GuestNice = float64(guestNice)
	}

	Info.Println("Return: CPU=", stat.CPU, " user=", stat.User, " Nice=", stat.Nice)
	return nil
}

