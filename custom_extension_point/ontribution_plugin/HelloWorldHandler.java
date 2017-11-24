package com.highgo.admin.migrator.handlers;

import java.util.Collection;

import org.eclipse.jface.dialogs.MessageDialog;
import org.eclipse.ui.IWorkbenchPart;
import org.eclipse.ui.IWorkbenchWindow;
import org.jkiss.dbeaver.DBException;
import org.jkiss.dbeaver.model.struct.DBSObject;
import org.jkiss.dbeaver.tools.IExternalTool;

public class HelloWorldHandler implements IExternalTool{

	@Override
	public void execute(IWorkbenchWindow window, 
			IWorkbenchPart activePart, Collection<DBSObject> objects)
			throws DBException{
		
		MessageDialog.openInformation(null, "E4 Information Dialog", 
				"Hello world from a pure Eclipse 4 plug-in");		
	}

}
